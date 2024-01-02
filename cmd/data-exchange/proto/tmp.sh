goass -cmd=generate -opt=unittest -file=../main.go -func=ProtoFoo -types=*data_exchange.RobotsFightData
goass -cmd=generate -opt=benchmark -file=../main.go -func=ProtoFoo -types=*data_exchange.RobotsFightData

goass -cmd=generate -opt=unittest -file=../main.go -func=JsonFoo -types=*data_exchange.RobotsFightData
goass -cmd=generate -opt=benchmark -file=../main.go -func=JsonFoo -types=*data_exchange.RobotsFightData

goass -cmd=generate -opt=unittest -file=../main.go -func=GobFoo -types=*data_exchange.RobotsFightData
goass -cmd=generate -opt=benchmark -file=../main.go -func=GobFoo -types=*data_exchange.RobotsFightData
