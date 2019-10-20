package utils

import "common"
import "config"

// GET
func PullObj(key string) (interface{}) {
	data := common.DeviceInfo_t { Entity:-1, Instant:-1, Key:key, ValueType:config.TYPE_OBJECT, Value: ""}

	msg := TalkToDao(config.GET_OBJ_BY_KEY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	if obj.ValueType != config.TYPE_RSP_ERROR {
		val := (obj.Value)
		return val
	}
	return nil
}
func PushObj(key string, val interface{}) (bool){
	data := common.DeviceInfo_t { Entity:-1, Instant:-1, Key:key, ValueType:config.TYPE_OBJECT, Value:val }
	msg := TalkToDao(config.SET_OBJ_BY_KEY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	if obj.ValueType == config.TYPE_RSP_OK {
		return true
	}
	return false
}

// GET
func PullObjDeviceMaxTemp(entity int, instant int) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_MAXTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PullObjDeviceAbsTemp(entity int, instant int) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_ABSTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PullObjListDeviceAbsTemp(entity int) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_LIST_ABSTEMP, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
func PullObjDeviceRelTemp(entity int, instant int) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_RELTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PullObjListDeviceRelTemp(entity int) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_TEMPERATURE, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_LIST_RELTEMP, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
// SET
func PushObjDeviceMaxTemp(entity int, instant int, value float64) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:value }
	msg := TalkToDao(config.SET_DEVICE_MAXTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PushObjDeviceAbsTemp(entity int, instant int, value float64) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:value }
	msg := TalkToDao(config.SET_DEVICE_ABSTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PushObjDeviceRelTemp(entity int, instant int, value float64) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_TEMPERATURE, Value:value }
	msg := TalkToDao(config.SET_DEVICE_RELTEMP, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// GET
func PullObjDeviceAveragePower(entity int, instant int) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_AVERAGEPOWER, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_AVERAGEPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PullObjListDeviceAveragePower(entity int) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_AVERAGEPOWER, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_LIST_AVERAGEPOWER, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
// SET
func PushObjDeviceAveragePower(entity int, instant int, value float64) (common.DeviceInfo_t){
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_AVERAGEPOWER, Value:value }
	msg := TalkToDao(config.SET_DEVICE_AVERAGEPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
// SET
func PushObjExpectFanDutyByTemperature(entity int, instant int, value float64) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_FANDUTY_TEMPERATURE, Value:value }
	msg := TalkToDao(config.SET_EXPECT_FAN_DUTY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}

func PushObjExpectFanDutyByPower(entity int, instant int, value float64) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_FANDUTY_POWER, Value:value }
	msg := TalkToDao(config.SET_EXPECT_FAN_DUTY, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
//GET
func PullObjDeviceMaxPower(entity int, instant int) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_MAXPOWER, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_MAXPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
func PullObjListDeviceMaxPower(entity int) (map[string]common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:-1, ValueType:config.TYPE_MAXPOWER, Value:0 }
	msg := TalkToDao(config.GET_DEVICE_LIST_MAXPOWER, data)
	obj := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj
}
//SET
func PushObjDeviceMaxPower(entity int, instant int, value float64) (common.DeviceInfo_t) {
	data := common.DeviceInfo_t { Entity:entity, Instant:instant, ValueType:config.TYPE_MAXPOWER, Value:value }
	msg := TalkToDao(config.SET_DEVICE_MAXPOWER, data)
	obj := (msg.Data).(common.DeviceInfo_t)
	return obj
}
//
func GetMapExpectFanDuty() (map[string]common.DeviceInfo_t) {
	obj := common.DeviceInfo_t { Entity: 0, Instant: 0, ValueType: "string", Value: "temp_" }
	msg := TalkToDao(config.GET_ALL_EXPECT_FAN_DUTY, obj)
	obj_map := (msg.Data).(map[string]common.DeviceInfo_t)
	return obj_map
}
