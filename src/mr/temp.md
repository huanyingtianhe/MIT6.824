1. worker注册信息到master
2. master分配一个map任务到一个空闲的worker
3. worker完成任务后响应给master
4. master等map任务列表全部完成后分配reduce任务
5. 等待worker完成所有的reduce任务

> master维持worker机器信息，存储任务的类型、状态