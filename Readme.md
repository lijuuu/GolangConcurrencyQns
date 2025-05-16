
---

###  Go Concurrency Questions

#### Question 1: Job Queue Dispatcher
**Problem**:
Simulate a job queue system where a dispatcher goroutine sends jobs ("Job1", "Job2", ...) to a pool of 3 worker goroutines, which print the job ID. Run for 5 seconds, ensuring jobs are distributed round-robin using a single buffered channel.

**Constraints**:
- Use one buffered channel (capacity 5) for job distribution.
- No `time.Sleep` or `for` loops for sending jobs.
- Use `Context` for the 5-second timeout.
- Ensure even job distribution across workers.
- No deadlocks or job loss.

**Expected Output**:
```
Worker1: Job1
Worker2: Job2
Worker3: Job3
Worker1: Job4
...
Done
```

[View Answer](./answers/qn1.go)

**Key Challenge**:
Implementing round-robin distribution with a single channel while ensuring all workers are utilized.

---

#### Question 2: Event Bus Broadcaster
**Problem**:
Create an event bus where a publisher goroutine sends events ("Event1", "Event2", ...) to 4 subscriber goroutines, each printing the event. Run for 6 seconds, using buffered channels to broadcast events to all subscribers.

**Constraints**:
- Use one buffered channel per subscriber (capacity 2).
- No `time.Sleep` or `for` loops for sending events.
- Use `context` for the 6-second timeout.
- Ensure all subscribers receive every event.
- No deadlocks or missed events.

**Expected Output**:
```
Subscriber1: Event1
Subscriber2: Event1
Subscriber3: Event1
Subscriber4: Event1
...
Done
```

**Key Challenge**:
Broadcasting events to multiple subscribers without losing synchronization or events.

---

#### Question 3: Database Connection Pool
**Problem**:
Simulate a database connection pool with 3 connections. Client goroutines (5 total) request connections, print "ClientX using ConnY", and release them after use. Run for 8 seconds, using a buffered channel to manage connection allocation.

**Constraints**:
- Use one buffered channel (capacity 3) for connections.
- No `time.Sleep` or `for` loops for requesting/releasing connections.
- Use `context` for the 8-second timeout.
- Ensure no client starves (all get at least one connection).
- No deadlocks or connection leaks.

**Expected Output**:
```
Client1 using Conn1
Client2 using Conn2
Client3 using Conn3
Client4 using Conn1
...
Done
```

**Key Challenge**:
Managing a limited resource pool with concurrent clients and ensuring fair access.

---

#### Question 4: Circuit Breaker
**Problem**:
Implement a circuit breaker that processes requests ("Req1", "Req2", ...) for 10 seconds. If 3 consecutive requests "fail" (simulated by `rand.Intn(2)`), switch to "Open" state and reject requests for 2 seconds, then retry. Print request status or state changes.

**Constraints**:
- Use one buffered channel (capacity 5) for requests.
- No `time.Sleep` or `for` loops for sending requests.
- Use `context` for the 10-second timeout.
- Use `time.After` for the 2-second open state.
- No deadlocks or state corruption.

**Expected Output**:
```
Processed Req1
Processed Req2
Circuit Open
Rejected Req3
Circuit Closed
...
Done
```

**Key Challenge**:
Managing state transitions with concurrent requests and timed recovery.

---

#### Question 5: Load Balancer
**Problem**:
Simulate a load balancer distributing requests ("Req1", "Req2", ...) to 3 backend servers, which print "ServerX: ReqY". Run for 7 seconds, using buffered channels to send requests to the least-loaded server (based on pending requests).

**Constraints**:
- Use one buffered channel per server (capacity 3).
- No `time.Sleep` or `for` loops for sending requests.
- Use `context` for the 7-second timeout.
- Track server load for balanced distribution.
- No deadlocks or request drops.

**Expected Output**:
```
Server1: Req1
Server2: Req2
Server3: Req3
Server1: Req4
...
Done
```

**Key Challenge**:
Dynamically selecting the least-loaded server while handling concurrent requests.

---

#### Question 6: Rate Limiter
**Problem**:
Create a rate limiter that allows 5 requests per second, printing "RequestX" for each. Run for 10 seconds, rejecting excess requests with "Rate Limited". Use a buffered channel and timer to enforce the rate.

**Constraints**:
- Use one buffered channel (capacity 5) for requests.
- No `time.Sleep` or `for` loops for sending requests.
- Use `context` for the 10-second timeout.
- Use `time.Tick` or `time.After` for timing.
- No deadlocks or timing drift.

**Expected Output**:
```
Request1
Request2
Rate Limited
Request3
...
Done
```

**Key Challenge**:
Enforcing a precise rate limit with concurrent requests and clean rejection handling.

---

#### Question 7: File Downloader Pipeline
**Problem**:
Simulate a file downloader with 3 stages: fetch URLs ("url1", "url2", ...), download content (print "Downloaded urlX"), and save (print "Saved urlX"). Run for 8 seconds, using buffered channels to pass data between stages.

**Constraints**:
- Use two buffered channels (capacity 4) for fetch-to-download and download-to-save.
- No `time.Sleep` or `for` loops for sending data.
- Use `context` for the 8-second timeout.
- Ensure all URLs are processed in order.
- No deadlocks or data loss.

**Expected Output**:
```
Downloaded url1
Saved url1
Downloaded url2
Saved url2
...
Done
```

**Key Challenge**:
Maintaining order in a multi-stage pipeline with concurrent processing.

---

#### Question 8: Voting System
**Problem**:
Implement a voting system where 5 voter goroutines submit votes ("Yes" or "No", chosen randomly) to a collector goroutine, which prints the vote count every 2 seconds for 10 seconds. Use a buffered channel for votes.

**Constraints**:
- Use one buffered channel (capacity 10) for votes.
- No `time.Sleep` or `for` loops for sending votes.
- Use `context` for the 10-second timeout.
- Print counts like "Yes: 3, No: 2" every 2 seconds.
- No deadlocks or vote loss.

**Expected Output**:
```
Yes: 2, No: 3
Yes: 5, No: 4
...
Done
```

**Key Challenge**:
Aggregating concurrent votes with periodic reporting and clean shutdown.

---

#### Question 9: Resource Cleanup Manager
**Problem**:
Simulate a system where 4 goroutines allocate resources ("Res1", "Res2", ...) and print "Allocated ResX". A cleanup goroutine deallocates them after 3 seconds, printing "Cleaned ResX". Run for 12 seconds, using buffered channels.

**Constraints**:
- Use one buffered channel (capacity 5) for allocation.
- No `time.Sleep` or `for` loops for sending resources.
- Use `context` for the 12-second timeout.
- Use `time.After` for the 3-second cleanup delay.
- No deadlocks or resource leaks.

**Expected Output**:
```
Allocated Res1
Allocated Res2
Cleaned Res1
Allocated Res3
...
Done
```

**Key Challenge**:
Coordinating allocation and timed cleanup without losing resources.

---

#### Question 10: Pub-Sub with Filters
**Problem**:
Create a pub-sub system where a publisher sends messages ("Msg1", "Msg2", ...) to 3 subscribers, each with a filter (e.g., Sub1 wants odd-numbered messages, Sub2 wants even, Sub3 wants all). Run for 6 seconds, using buffered channels.

**Constraints**:
- Use one buffered channel per subscriber (capacity 2).
- No `time.Sleep` or `for` loops for sending messages.
- Use `context` for the 6-second timeout.
- Ensure correct filtering (e.g., Sub1 gets Msg1, Msg3, ...).
- No deadlocks or missed messages.

**Expected Output**:
```
Sub1: Msg1
Sub2: Msg2
Sub3: Msg1
Sub3: Msg2
...
Done
```

**Key Challenge**:
Implementing filtered message delivery with concurrent subscribers.

---

#### Question 11: Task Retry Handler
**Problem**:
Simulate a task processor that retries failed tasks ("Task1", "Task2", ...) up to 3 times before printing "Failed TaskX". Run for 10 seconds, with failures simulated by `rand.Intn(2)`. Use a buffered channel for tasks.

**Constraints**:
- Use one buffered channel (capacity 5) for tasks.
- No `time.Sleep` or `for` loops for sending tasks.
- Use `context` for the 10-second timeout.
- Track retry counts per task.
- No deadlocks or task loss.

**Expected Output**:
```
Processed Task1
Retrying Task2
Failed Task2
...
Done
```

**Key Challenge**:
Managing retry state for concurrent tasks with failure simulation.

---

#### Question 12: Priority Queue Processor
**Problem**:
Implement a priority queue that processes tasks with priorities ("High1", "Low1", ...) for 8 seconds. High-priority tasks always run before low-priority ones. Use two buffered channels for high and low tasks.

**Constraints**:
- Use two buffered channels (capacity 3 each).
- No `time.Sleep` or `for` loops for sending tasks.
- Use `context` for the 8-second timeout.
- Ensure high-priority tasks are never starved.
- No deadlocks or task drops.

**Expected Output**:
```
High1
High2
Low1
High3
...
Done
```

**Key Challenge**:
Enforcing priority in a concurrent queue without starving low-priority tasks.

---

#### Question 13: Heartbeat Monitor
**Problem**:
Simulate a system where 3 services send heartbeats ("Service1 OK", ...) every 2 seconds to a monitor goroutine, which prints "All Services OK" if all heartbeats are received, or "ServiceX Missing" if any fail. Run for 10 seconds.

**Constraints**:
- Use one buffered channel per service (capacity 2).
- No `time.Sleep` or `for` loops for sending heartbeats.
- Use `context` for the 10-second timeout.
- Use `time.After` for heartbeat timing.
- No deadlocks or false reports.

**Expected Output**:
```
All Services OK
Service2 Missing
All Services OK
...
Done
```

**Key Challenge**:
Tracking periodic heartbeats with concurrent services and accurate status reporting.

---

#### Question 14: Cache Eviction System
**Problem**:
Simulate a cache where 4 goroutines add items ("Item1", "Item2", ...) and an evictor goroutine removes the oldest item every 3 seconds, printing "Evicted ItemX". Run for 12 seconds, using a buffered channel.

**Constraints**:
- Use one buffered channel (capacity 5) for items.
- No `time.Sleep` or `for` loops for adding items.
- Use `context` for the 12-second timeout.
- Use `time.After` for eviction timing.
- No deadlocks or item loss.

**Expected Output**:
```
Added Item1
Added Item2
Evicted Item1
Added Item3
...
Done
```

**Key Challenge**:
Maintaining a FIFO cache with concurrent additions and timed evictions.

---

#### Question 15: Barrier Synchronization
**Problem**:
Implement a barrier where 5 goroutines perform tasks ("Task1", "Task2", ...) and wait at a barrier before printing "Barrier Reached". Run for 10 seconds, releasing all goroutines every 2 seconds.

**Constraints**:
- Use one buffered channel (capacity 5) for barrier synchronization.
- No `time.Sleep` or `for` loops for sending tasks.
- Use `context` for the 10-second timeout.
- Use `time.After` for barrier timing.
- No deadlocks or missed synchronizations.

**Expected Output**:
```
Task1
Task2
Barrier Reached
Task3
...
Done
```

**Key Challenge**:
Synchronizing multiple goroutines at a timed barrier with concurrent tasks.

---

#### Question 16: Throttled API Caller
**Problem**:
Simulate an API caller that sends requests ("Req1", "Req2", ...) to a server with a limit of 3 concurrent requests. Print "ReqX Sent" and "ReqX Done" for each. Run for 8 seconds, using a buffered channel.

**Constraints**:
- Use one buffered channel (capacity 3) for requests.
- No `time.Sleep` or `for` loops for sending requests.
- Use `context` for the 8-second timeout.
- Ensure exactly 3 requests are in flight at a time.
- No deadlocks or request overruns.

**Expected Output**:
```
Req1 Sent
Req2 Sent
Req3 Sent
Req1 Done
Req4 Sent
...
Done
```

**Key Challenge**:
Throttling concurrent requests with precise control over in-flight tasks.

---

#### Question 17: Log Aggregator
**Problem**:
Create a log aggregator where 4 goroutines generate logs ("Log1 from Src1", ...) and a collector goroutine prints aggregated logs every 2 seconds ("Src1: 3 logs, Src2: 2 logs"). Run for 10 seconds.

**Constraints**:
- Use one buffered channel (capacity 10) for logs.
- No `time.Sleep` or `for` loops for sending logs.
- Use `context` for the 10-second timeout.
- Use `time.After` for aggregation timing.
- No deadlocks or log loss.

**Expected Output**:
```
Src1: 2 logs, Src2: 1 log, Src3: 3 logs
Src1: 1 log, Src4: 2 logs
...
Done
```

**Key Challenge**:
Aggregating logs from multiple sources with periodic reporting.

---

#### Question 18: Worker Failure Recovery
**Problem**:
Simulate a system with 3 workers processing tasks ("Task1", "Task2", ...). If a worker "fails" (10% chance via `rand.Float64()`), a new worker takes over, printing "WorkerX Recovered". Run for 10 seconds.

**Constraints**:
- Use one buffered channel (capacity 5) for tasks.
- No `time.Sleep` or `for` loops for sending tasks.
- Use `context` for the 10-second timeout.
- Ensure seamless task processing despite failures.
- No deadlocks or task loss.

**Expected Output**:
```
Task1
Task2
Worker2 Recovered
Task3
...
Done
```

**Key Challenge**:
Dynamically replacing failed workers without interrupting task flow.

---

#### Question 19: Ordered Broadcast
**Problem**:
Implement a broadcast system where a sender goroutine sends messages ("Msg1", "Msg2", ...) to 3 receivers, which print "ReceiverX: MsgY" in a fixed order (Receiver1, Receiver2, Receiver3). Run for 8 seconds.

**Constraints**:
- Use one buffered channel per receiver (capacity 2).
- No `time.Sleep` or `for` loops for sending messages.
- Use `context` for the 8-second timeout.
- Ensure strict receiver order for each message.
- No deadlocks or message skips.

**Expected Output**:
```
Receiver1: Msg1
Receiver2: Msg1
Receiver3: Msg1
Receiver1: Msg2
...
Done
```

**Key Challenge**:
Enforcing ordered delivery across concurrent receivers for each broadcast.

---

#### Question 20: Dynamic Task Splitter
**Problem**:
Simulate a task splitter where a generator sends tasks ("Task1", "Task2", ...) to 2–4 worker goroutines (chosen randomly at startup), which print "WorkerX: TaskY". Run for 10 seconds, splitting tasks evenly.

**Constraints**:
- Use one buffered channel per worker (capacity 3).
- No `time.Sleep` or `for` loops for sending tasks.
- Use `context` for the 10-second timeout.
- Randomly select 2–4 workers at startup.
- No deadlocks or uneven task splits.

**Expected Output**:
```
Worker1: Task1
Worker2: Task2
Worker3: Task3
Worker1: Task4
...
Done
```

**Key Challenge**:
Dynamically configuring workers and ensuring balanced task distribution.
