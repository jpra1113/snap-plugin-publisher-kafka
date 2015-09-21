// +build unit

package kafka

import (
	"testing"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	"github.com/intelsdi-x/pulse/core/ctypes"
	. "github.com/smartystreets/goconvey/convey"
)

func TestKafkaPlugin(t *testing.T) {
	Convey("Meta returns proper metadata", t, func() {
		meta := Meta()
		So(meta.Name, ShouldResemble, PluginName)
		So(meta.Version, ShouldResemble, PluginVersion)
		So(meta.Type, ShouldResemble, plugin.PublisherPluginType)
	})

	Convey("Create Kafka Publisher", t, func() {
		k := NewKafkaPublisher()
		Convey("so kafka publisher should not be nil", func() {
			So(k, ShouldNotBeNil)
		})
		Convey("so kafka publisher should be of kafka publisher type", func() {
			So(k, ShouldHaveSameTypeAs, &kafkaPublisher{})
		})
		Convey("k.GetConfigPolicy()", func() {
			configPolicy := k.GetConfigPolicy()
			Convey("So config policy should not be nil", func() {
				So(configPolicy, ShouldNotBeNil)
			})
			Convey("So config policy should be a cpolicy.ConfigPolicy type", func() {
				So(configPolicy, ShouldHaveSameTypeAs, cpolicy.ConfigPolicy{})
			})
			testConfig := make(map[string]ctypes.ConfigValue)
			testConfig["brokers"] = ctypes.ConfigValueStr{Value: "127.0.0.1:9092"}
			testConfig["topic"] = ctypes.ConfigValueStr{Value: "test"}
			cfg, errs := configPolicy.Get([]string{""}).Process(testConfig)
			Convey("So config policy should process testConfig and return a config", func() {
				So(cfg, ShouldNotBeNil)
			})
			Convey("So testConfig processing should return no errors", func() {
				So(errs.HasErrors(), ShouldBeFalse)
			})
		})
	})
}
