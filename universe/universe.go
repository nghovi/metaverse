package universe

type SelfIntroducer interface {
	SelfIntroduce()
}

type Speaker interface {
	Speak(sentence string)
}

type Eater interface {
	Eat()
}
