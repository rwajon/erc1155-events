package db

func Init() {
	Transaction.createIndexes()
	WatchList.createIndexes()
}
