#!/bin/bash

SRV_HOST="127.0.0.1:8088"
# AIC MAX Temperature
DATA='{"entity":3, "instant":1, "value":100.0}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/maxtemp
echo ""
DATA='{"entity":3, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/maxtemp
echo ""

#
DATA='{"entity":3, "instant":1, "value":52.12}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/set/device/abstemp
echo ""

DATA='{"entity":3, "instant":1}'
curl -d "$DATA" -H "Content-Type: application/json" -X POST http://$SRV_HOST/api/v1/hwmon/get/device/abstemp
echo ""
