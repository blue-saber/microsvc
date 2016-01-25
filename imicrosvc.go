package microsvc

type IMicroService interface {
	Run(addr string)
}
