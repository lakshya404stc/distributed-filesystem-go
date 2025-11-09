package p2p

type HandShakeFunc func(any) error

func NOPHandshake (any) error {return nil}