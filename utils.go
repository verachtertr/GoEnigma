package Enigma

func modulo(a,b int) int{
  temp := a % b
  if temp < 0 {
    return (temp + b)
  }
  return temp
}
