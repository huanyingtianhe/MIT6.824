1. worker注册信息到master
2. master分配一个map任务到一个空闲的worker
3. worker完成任务后响应给master
4. master等map任务列表全部完成后分配reduce任务
5. 等待worker完成所有的reduce任务
6. 所有的map和reduce任务都完成，master唤醒用户程序。用户程序的mapReduce调用才返回。

> master维持worker机器信息，存储每一个Map和Reduce任务(空闲、工作中或完成)的状态。

> Map任务完成时，Master接受到位置和大小的更新信息，然后把信息逐步递增地推送给正在工作的Reduce任务。