namespace go thrifttest.duck

struct Duck {
  1: optional string quack,
}

service Feeder {
  void feed(1:Duck duck),
}
