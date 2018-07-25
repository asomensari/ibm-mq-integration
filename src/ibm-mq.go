package main

import (
	"github.com/ibm-messaging/mq-golang/mqmetric"
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	log2 "github.com/newrelic/infra-integrations-sdk/log"
	"strings"
)

const (
	integrationName    = "com.expertservices.ibm-mq"
	integrationVersion = "0.1.0"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	QueueManager string `default:"QMA" help:"Queue Manager name"`

	UserName string `help:"MQ Username"`
	Password string `help:"Password"`

	ReplyQueue     string `default:"SYSTEM.DEFAULT.MODEL.QUEUE" help:"Reply Queue to collect data"`
	MonitoredQueue string `help:"Patterns of queues to monitor"`
}

var (
	args      argumentList
	eventData map[string]interface{}
	cc        mqmetric.ConnectionConfig
	log       log2.Logger
)

func main() {

	//remoteEntities := true
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	panicOnErr(err)

	//log.Logger().Infof("Starting IBM MQ")
	//fmt.Println(args.MonitoredQueue)

	i.Logger().Infof("Starting IBM MQ metrics exporter for JSON")
	//log.Infoln("Starting IBM MQ metrics exporter for JSON")

	cc.ClientMode = false
	if args.UserName != "" {
		cc.UserId = args.UserName
		cc.Password = args.Password
	}

	// Connect and open standard queues
	err = mqmetric.InitConnection(args.QueueManager, "SYSTEM.DEFAULT.MODEL.QUEUE", &cc)

	if err == nil {
		//i.Logger().Infof("Connected to queue manager %s ", string(args.QueueManager))
		defer mqmetric.EndConnection()
	}

	// What metrics can the queue manager provide? Find out, and
	// subscribe.
	if err == nil {
		err = mqmetric.DiscoverAndSubscribe(args.MonitoredQueue, true, "")
	}

	data := Collect(i)

	//for e := range data {
	//		fmt.Println(e)
	//}

	//if remoteEntities {
	//integrationWithRemoteEntities(i,data)
	//} else {
	integrationWithLocalEntity(i, data)
	//}
	panicOnErr(i.Publish())
}

//func integrationWithRemoteEntities(i *integration.Integration, data map[string]interface{}) {

// Create Entity, entities name must be unique
//	e1, err := i.Entity("redis-instance-1", "cache")
//	panicOnErr(err)

// Add an Event
//e1.AddEvent(event.New("restart", "status"))

// Add Inventory item
//e1.SetInventoryItem("redis-server", "version", "3.0.1")
//e1.SetInventoryItem("connection", "type", "redis-slave1.example.com")

// Add Metric
//	m1, err := e1.NewMetricSet("IBMMQSample")

//	populateMetrics(m1,data,defaultMetrics);

//}

func integrationWithLocalEntity(i *integration.Integration, data map[string]interface{}) {

	// This integration will show any data attached to the host where the agent is running
	//localEntity := i.LocalEntity()
	//localEntity.AddEvent(event.New("restart", "status"))

	// Add Inventory item
	//localEntity.SetInventoryItem("redis-server", "version", "3.0.1")

	localEntity := i.LocalEntity()

	// Add Metric
	m, err := localEntity.NewMetricSet("IBMMQSample")
	panicOnErr(err)

	populateMetrics(m, data, defaultMetrics)
	//m.SetMetric("keys", 4000, metric.GAUGE)
	//m.SetMetric("server-type", "master", metric.ATTRIBUTE)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Collect(i *integration.Integration) map[string]interface{} {

	var err error
	eventData = make(map[string]interface{})

	i.Logger().Infof("IBM MQ JSON collector started")

	// Clear out everything we know so far. In particular, replace
	// the map of values for each object so the collection starts
	// clean.

	if args.QueueManager == "" {
		//err = err.Error("No Queue Manager Defined")
		panicOnErr(err)
	}

	eventData["queueManager"] = string(args.QueueManager)
	if args.MonitoredQueue != "" {
		eventData["monitoredQueue"] = args.MonitoredQueue
	}
	err = mqmetric.InitConnection(args.QueueManager, "SYSTEM.DEFAULT.MODEL.QUEUE", &cc)
	if err == nil {
		i.Logger().Infof("Connected to queue manager ", args.QueueManager)
		defer mqmetric.EndConnection()
	}

	// What metrics can the queue manager provide? Find out, and
	// subscribe.
	if err == nil {
		err = mqmetric.DiscoverAndSubscribe(args.MonitoredQueue, true, "")
	}

	for _, cl := range mqmetric.Metrics.Classes {
		for _, ty := range cl.Types {
			for _, elem := range ty.Elements {
				elem.Values = make(map[string]int64)
			}
		}
	}

	// Process all the publications that have arrived
	mqmetric.ProcessPublications()

	// Have now processed all of the publications, and all the MQ-owned
	// value fields and maps have been updated.
	//
	// Now need to set all of the real items with the correct values
	for _, cl := range mqmetric.Metrics.Classes {
		for _, ty := range cl.Types {
			for _, elem := range ty.Elements {
				for key, value := range elem.Values {
					f := mqmetric.Normalise(elem, key, value)
					tags := map[string]string{
						"qmgr": args.QueueManager,
					}

					if key != mqmetric.QMgrMapKey {
						tags["object"] = key
					}
					printPoint(elem.MetricName, float32(f), tags)

				}
			}
		}
	}

	if err != nil {
		i.Logger().Errorf(" Error Posting %s/n", err)
	}

	return eventData
}

func printPoint(metric string, val float32, tags map[string]string) {

	if q, ok := tags["object"]; ok {
		eventData["queue"] = q
	}

	eventData[fixup(metric)] = val

	return
}

func fixup(s1 string) string {
	// Another reformatting of the metric name - this one converts
	// something like queue_avoided_bytes into queueAvoidedBytes
	s2 := ""
	c := ""
	nextCaseUpper := false

	for i := 0; i < len(s1); i++ {
		if s1[i] != '_' {
			if nextCaseUpper {
				c = strings.ToUpper(s1[i : i+1])
				nextCaseUpper = false
			} else {
				c = strings.ToLower(s1[i : i+1])
			}
			s2 += c
		} else {
			nextCaseUpper = true
		}

	}
	return s2
}
