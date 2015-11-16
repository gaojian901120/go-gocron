
type Cron struct {
    entries  []*Entry
    stop     chan struct{}   // 控制 Cron 实例暂停
    <span id="12_nwp" style="width: auto; height: auto; float: none;"><a id="12_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=add&k0=add&kdi0=0&luki=4&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="12" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">add</span></a></span>      chan *Entry     // 当 Cron 已经运行了，增加新的 Entity 是通过 add 这个 channel 实现的
    snapshot chan []*Entry   // 获取当前所有 entity 的快照
    running  bool            // 当已经运行时为true；否则为false
}
type Entry struct {
    // The schedule on which this <span id="9_nwp" style="width: auto; height: auto; float: none;"><a id="9_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=job&k0=job&kdi0=0&luki=10&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="9" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">job</span></a></span> should be run.
    // 负责调度当前 Entity 中的 Job 执行
    Schedule Schedule

    // The <span id="10_nwp" style="width: auto; height: auto; float: none;"><a id="10_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=next&k0=next&kdi0=0&luki=5&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="10" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">next</span></a></span> time the job will run. This is the <span id="11_nwp" style="width: auto; height: auto; float: none;"><a id="11_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=zero&k0=zero&kdi0=0&luki=2&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="11" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">zero</span></a></span> time if Cron has not been
    // started or this entry's schedule is unsatisfiable
    // Job 下一次执行的时间
    Next time.Time

    // The last time this job was run. This is the zero time if the job has never
    // been run.
    // 上一次执行时间
    Prev time.Time

    // The Job to run.
    // 要执行的 Job
    Job Job
}
type Job interface {
    Run()
}
type FuncJob func()
func (f FuncJob) Run() { f() }
type Schedule interface {
    // Return the <span id="6_nwp" style="width: auto; height: auto; float: none;"><a id="6_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=next&k0=next&kdi0=0&luki=5&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="6" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">next</span></a></span> activation time, later than the given time.
    // Next is invoked initially, and then each time the <span id="7_nwp" style="width: auto; height: auto; float: none;"><a id="7_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=job&k0=job&kdi0=0&luki=10&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="7" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">job</span></a></span> is run.
    // 返回同一 Entity 中的 Job 下一次执行的时间
    Next(time.Time) time.Time
}

type SpecSchedule struct {
    Second, Minute, Hour, Dom, Month, Dow uint64
}

type ConstantDelaySchedule struct {
    Delay time.Duration // 循环的时间间隔
}

constDelaySchedule := Every(5e9)
func New() *Cron {
    return &Cron{
        entries:  nil,
        <span id="4_nwp" style="width: auto; height: auto; float: none;"><a id="4_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=add&k0=add&kdi0=0&luki=4&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="4" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">add</span></a></span>:      make(chan *Entry),
        stop:     make(chan struct{}),
        snapshot: make(chan []*Entry),
        running:  false,
    }
}

func Parse(spec string) (_ Schedule, err error)


// 将 <span id="1_nwp" style="width: auto; height: auto; float: none;"><a id="1_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=job&k0=job&kdi0=0&luki=10&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="1" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">job</span></a></span> 加入 Cron 中
// 如上所述，该方法只是简单的通过 FuncJob 类型强制转换 cmd，然后调用 AddJob 方法
func (c *Cron) AddFunc(spec string, cmd func()) error

// 将 job 加入 Cron 中
// 通过 Parse <span id="2_nwp" style="width: auto; height: auto; float: none;"><a id="2_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=%BA%AF%CA%FD&k0=%BA%AF%CA%FD&kdi0=0&luki=1&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="2" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">函数</span></a></span>解析 cron 表达式 spec 的到调度器实例(Schedule)，之后调用 c.Schedule 方法
func (c *Cron) AddJob(spec string, cmd Job) error

// 获取当前 Cron 总所有 Entities 的快照
func (c *Cron) Entries() []*Entry

// 通过两个参数实例化一个 Entity，然后加入当前 Cron 中
// 注意：如果当前 Cron 未运行，则直接将该 entity 加入 Cron 中；
// 否则，通过 <span id="3_nwp" style="width: auto; height: auto; float: none;"><a id="3_nwl" href="http://cpro.baidu.com/cpro/ui/uijs.php?adclass=0&app_id=0&c=news&cf=1001&ch=0&di=128&fv=18&is_app=0&jk=ffe7e54aca7ba4d3&k=add&k0=add&kdi0=0&luki=4&mcpm=0&n=10&p=baidu&q=74042097_cpr&rb=0&rs=1&seller_id=1&sid=d3a47bca4ae5e7ff&ssp2=1&stid=0&t=tpclicked3_hc&td=1989498&tu=u1989498&u=http%3A%2F%2Fblog%2Estudygolang%2Ecom%2F2014%2F02%2Fgo%5Fcrontab%2F&urlid=0" target="_blank" mpid="3" style="text-decoration: none;"><span style="color:#0000ff;font-size:13.92px;width:auto;height:auto;float:none;">add</span></a></span> 这个成员 channel 将 entity 加入正在运行的 Cron 中
func (c *Cron) Schedule(schedule Schedule, cmd Job)

// 新启动一个 goroutine 运行当前 Cron
func (c *Cron) Start()

// 通过给 stop 成员发送一个 struct{}{} 来停止当前 Cron，同时将 running 置为 false
// 从这里知道，stop 只是通知 Cron 停止，因此往 channel 发一个值即可，而不关心值是多少
// 所以，成员 stop 定义为空 struct
func (c *Cron) Stop()


package main

import (
    "github.com/robfig/cron"
    "log"
)

func main() {
    i := 0
    c := cron.New()
    spec := "*/5 * * * * ?"
    c.AddFunc(spec, func() {
        i++
        log.Println("cron running:", i)
    })
    c.Start()

    select{}
}
