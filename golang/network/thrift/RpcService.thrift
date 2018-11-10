namespace go demo.rpc
namespace java demo.rpc

// ²âÊÔ·þÎñ
service RpcService {
	list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
}
