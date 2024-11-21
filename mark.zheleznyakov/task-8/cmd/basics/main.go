package main

func concatFirstTwo(a []byte, b []byte) []byte {
  return append(a[:2], b[:2]...)
}

func main() {
  concatResult := concatFirstTwo([]byte("occured"), []byte("toaster"))
  println(string(concatResult))
}
