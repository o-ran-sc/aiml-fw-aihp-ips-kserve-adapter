.. This work is licensed under a Creative Commons Attribution 4.0 International License.
.. http://creativecommons.org/licenses/by/4.0

.. Copyright (c) 2023 Samsung Electronics Co., Ltd. All Rights Reserved.

User-Guide
==========

.. contents::
   :depth: 3
   :local:


Overview
--------
Kserve Adapter works with the AIML Framework and is used to deploy, delete and update AI/ML models on Kserve.


Steps to build and run kserve adapter
-------------------------------------

Prerequisites

#. Install go
#. Install make


Steps

.. code:: bash

        git clone "https://gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter"
        cd kserve-adapter

| Update ENV variables in Makefile under run section.
| Update :file:`pkg/helm/data/sample_config.json` with model url. This can be obtained from AIMFW dashboard(Training Jobs-> Training Job status -> Select Info for a training job -> Model URL)
| Execute below commands
        
.. code:: bash

        make build
        make run


Steps to run kserve adapter using AIMLFW deployment scripts
-----------------------------------------------------------

Follow the steps in this link: `AIMLFW installation guide <https://docs.o-ran-sc.org/projects/o-ran-sc-aiml-fw-aimlfw-dep/en/latest/installation-guide.html>`__

Demo steps
----------

Prerequisites

#. Install chart museum
#. Build ricdms binary


Steps for the demo

#. Run ric dms
   
   .. code:: bash

        export RIC_DMS_CONFIG_FILE=$(pwd)/config/config-test.yaml
        ./ricdms

#. Run kserve adapter

   Create namespace called ricips

   .. code:: bash

        kubectl create ns ricips

|  Update ENV variables in Makefile under run section.
|  Update :file: `pkg/helm/data/sample_config.json` with model url. This can be obtained from AIMFW dashboard(Training Jobs-> Training Job status -> Select Info for a training job -> Model URL)
|  Execute below commands

   .. code:: bash

        make build
        make run

#. Generating and upload helm package

   .. code:: bash

        curl --request POST --url 'http://127.0.0.1:10000/v1/ips/preparation?configfile=pkg/helm/data/sample_config.json&schemafile=pkg/helm/data/sample_schema.json'

#. Check uploaded charts

   .. code:: bash

        curl http://127.0.0.1:8080/api/charts

#. Deploying the model

   .. code:: bash

        curl --request POST --url 'http://127.0.0.1:10000/v1/ips?name=inference-service&version=1.0.0'

#. Check deployed Inference service

   .. code:: bash

        kubectl get InferenceService -n ricips

#. Perform predictions

   Use below command to obtain Ingress port for Kserve.

   .. code:: bash

        kubectl get svc istio-ingressgateway -n istio-system

  
   Obtain nodeport corresponding to port 80.
   In the below example, the port is 31206.

   .. code:: bash

           NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                                                                      AGE
           istio-ingressgateway   LoadBalancer   10.105.222.242   <pending>     15021:31423/TCP,80:31206/TCP,443:32145/TCP,31400:32338/TCP,15443:31846/TCP   4h15m
   
   
  Create file predict_inference.sh with below contents:

   .. code:: bash

        model_name=sample-xapp
        curl -v -H "Host: $model_name.ricips.example.com" http://<VM IP>:<Ingress port for Kserve>/v1/models/$model_name:predict -d @./input_qoe.json

  Update the VM IP and the Ingress port for Kserve above. 

  Create file input_qoe.json with below contents:

   .. code:: bash

        {"signature_name": "serving_default", "instances": [[[2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56],
                [2.56, 2.56]]]}

  Use command below to trigger predictions

  .. code:: bash

        source predict_inference.sh

