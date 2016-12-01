package pbg

PokèmonDB interface {

}

PokèmonDBInterface interface {
	Get(id int) (Pokèmon, error)
}

MoveDB interface {

}

MoveDBInterface interface {
	Get(id int) (Move, error)
}

TrainerDB interface {

}

TrainerDBInterface interface {
	Get(id string) (Trainer, error)
	Add(user, password string) (Trainer, error)
	Delete(Trainer)
}

SessionDB interface {
	Supply() SessionDBInterface
	Retrieve(SessionDBInterface)
}

SessionDBInterface interface {
	Get(id string) (Session, error)
	Add(trainer Trainer) (Session, error)
	Delete(session Session)
}