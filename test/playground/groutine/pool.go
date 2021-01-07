package groutine

// 开启协程池

// 1、需要有workerChan
// 2、需要有resultChan

// resultChan用来收数据
// workerChan用来收任务

// 创建池其就是就是开协程
// 每个协程从workerChan收初始数据 进行运算 发给resultChan

// 主进程用来给workerChan传数据

/*func createPool(num int, i)  {

}
*/
