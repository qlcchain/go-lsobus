package utils

import pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

func GenerateWork(hash pkg.Hash) pkg.Work {
	var w pkg.Work
	worker, _ := pkg.NewWorker(w, hash)
	return worker.NewWork()
}
