package migration

// Migration contract
type Migration interface {
	Run() error
}
