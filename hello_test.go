package min

import "testing"

func TestMin(t *testing.T) {
   a := 1
   b := 2
   res := Min(a, b)
   if res == false {
       t.Errorf("test error, %d, %d", a, b)
   }
}

func TestMinErr(t *testing.T) {
   a := 2
   b := 3
   res := Min(a, b)
   if res == false {
       t.Errorf("test error, %d, %d", a, b)
   }
}
