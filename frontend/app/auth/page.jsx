'use client'

import Login from "@/components/Login"
import Signup from "@/components/Signup"
import { useState } from "react"

export default function Auth() {

  const [isLogin, setIsLogin] = useState(true)

  return (
    <div>
      {isLogin ?
        <Login setIsLogin={setIsLogin} />
        :
        <Signup setIsLogin={setIsLogin} />
      }
    </div>
  )
}