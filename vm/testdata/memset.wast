(module
 (type $FUNCSIG$viii (func (param i32 i32 i32)))
 (type $FUNCSIG$iiii (func (param i32 i32 i32) (result i32)))
 (import "env" "cos_assert" (func $cos_assert (param i32 i32 i32)))
 (import "env" "memset" (func $memset (param i32 i32 i32) (result i32)))
 (table 0 anyfunc)
 (memory $0 1)
 (data (i32.const 4) " @\00\00")
 (data (i32.const 16) "hello\00")
 (export "memory" (memory $0))
 (export "main" (func $main))
 (func $main (result i32)
  (local $0 i32)
  (i32.store offset=4
   (i32.const 0)
   (tee_local $0
    (i32.sub
     (i32.load offset=4
      (i32.const 0)
     )
     (i32.const 128)
    )
   )
  )
  (drop
   (call $memset
    (i32.add
     (get_local $0)
     (i32.const 16)
    )
    (i32.const 0)
    (i32.const 100)
   )
  )
  (i32.store16
   (i32.add
    (get_local $0)
    (i32.const 12)
   )
   (i32.load16_u offset=20 align=1
    (i32.const 0)
   )
  )
  (i32.store16 offset=16
   (get_local $0)
   (i32.const 770)
  )
  (i32.store offset=8
   (get_local $0)
   (i32.load offset=16 align=1
    (i32.const 0)
   )
  )
  (call $cos_assert
   (i32.eq
    (i32.load8_s offset=16
     (get_local $0)
    )
    (i32.const 2)
   )
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (i32.const 20)
  )
  (call $cos_assert
   (i32.eq
    (i32.load8_s offset=17
     (get_local $0)
    )
    (i32.const 3)
   )
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (i32.const 20)
  )
  (call $cos_assert
   (i32.eqz
    (i32.load8_s offset=115
     (get_local $0)
    )
   )
   (i32.add
    (get_local $0)
    (i32.const 8)
   )
   (i32.const 20)
  )
  (i32.store offset=4
   (i32.const 0)
   (i32.add
    (get_local $0)
    (i32.const 128)
   )
  )
  (i32.const 0)
 )
)
