package linkedin

import (
    "sync"
    "time"

    "iTools/crawler"
)

const (
    jWorker = 50
)

func mainWork(wg *sync.WaitGroup) {
    for job := range jobsCh {
        link := crawler.GetDocLink(job.Work)

        // link不為空
        if link != "" {
            doneCh <- JobDone{job, link}
        }
    }
    wg.Done()
}

func JobQueene(workLists []string) []string {
    workMax = len(workLists)

    go addList(workLists)
    go jodDone(finishCh)
    createWorkerPool(jWorker)

    <-finishCh

    jobsCh = make(chan Job, jWorker)
    doneCh = make(chan JobDone, jWorker)

    return workResults
}

var (
    jprocess = 0
    start    = time.Now()

    jobsCh   = make(chan Job, jWorker)
    doneCh   = make(chan JobDone, jWorker)
    finishCh = make(chan bool)

    workResults []string
    workMax     int
)

func addList(workLists []string) {
    for i, workList := range workLists {
        jobsCh <- Job{i, workList}
    }
    close(jobsCh)
}

func jodDone(finishCh chan bool) {
    bubbleSortJob := func(arr []JobDone) {
        size := len(arr)
        for i := 0; i < size; i++ {
            for j := 1; j < size; j++ {
                if arr[j].Job.Id < arr[j-1].Job.Id {
                    arr[j], arr[j-1] = arr[j-1], arr[j]
                }
            }
        }
    }

    var wr []JobDone
    for r := range doneCh {
        wr = append(wr, r)
    }
    bubbleSortJob(wr)

    workResults = []string{}

    for _, r := range wr {
        workResults = append(workResults, r.Work)
    }

    finishCh <- true
}

func createWorkerPool(jWorker int) {
    var wg sync.WaitGroup
    for i := 0; i < jWorker; i++ {
        wg.Add(1)
        go mainWork(&wg)
    }
    wg.Wait()
    close(doneCh)
}

type Job struct {
    Id   int
    Work string
}

type JobDone struct {
    Job  Job
    Work string
}
