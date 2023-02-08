<H1>Simple chatting service.</H1>
Chat service with grpc API and streaming.
User-to-user and group chatting.

Server and Client located under ./cmd

Start server with command
```
make run-server
```

Then open the second terminal and start client with command
```
make run-client
```

Check logs to look similar to the example below

For server:
```
2023/02/08 22:32:14 server listening at [::]:13111
2023/02/08 22:32:20 client id 222 connected
2023/02/08 22:32:20 client id 111 connected
2023/02/08 22:32:20 client id 333 connected
2023/02/08 22:32:25 failed to register listener: RegisterListener: client is already present in group
2023/02/08 22:32:25 sendGroupMessages: group 222 not found, skip to users: GetListeners: listeners not found
2023/02/08 22:32:25 sendUserMessage: user chat10 not found: SendMessage: client's data was not found
2023/02/08 22:32:25 Client id 222 disconnected
2023/02/08 22:32:25 UnregisterListenerFromAllGroups: client is not present in group chat10
2023/02/08 22:32:25 Client id 111 disconnected
2023/02/08 22:32:25 Client id 333 disconnected

```

For client:
```
2023/02/08 22:32:25 CreateGroupChat: in chat_name:"chat10" username:"111", out , err <nil>
2023/02/08 22:32:25 CreateGroupChat: in chat_name:"chat10" username:"222", out , err <nil>
2023/02/08 22:32:25 CreateGroupChat: in chat_name:"chat10" username:"222", out <nil>, err rpc error: code = Internal desc = internal error
2023/02/08 22:32:25 CreateGroupChat: in chat_name:"chat10" username:"333", out , err <nil>
2023/02/08 22:32:25 LeftGroupChat: in chat_name:"chat10" username:"333", out , err <nil>
2023/02/08 22:32:25 SendMessage: in data:{sender_id:"111" receiver_id:"222" payload:"hello from 111 to 222"}, out , err <nil>
2023/02/08 22:32:25 Message received from stream 222: sender_id:"111" receiver_id:"222" payload:"hello from 111 to 222"
2023/02/08 22:32:25 SendMessage: in data:{sender_id:"111" receiver_id:"chat10" payload:"hello from 111 to chan10"}, out , err <nil>
2023/02/08 22:32:25 Message received from stream 111: sender_id:"chat10" receiver_id:"111" payload:"hello from 111 to chan10"
2023/02/08 22:32:25 Message received from stream 222: sender_id:"chat10" receiver_id:"222" payload:"hello from 111 to chan10"
2023/02/08 22:32:25 SendMessage: in , out groups:{group_id:"chat10" usernames:"111" usernames:"222"}, err <nil>
2023/02/08 22:32:25 client actions finished. Shutting down

```

To re-generate protobuf schema type command
```
make generate  
```
