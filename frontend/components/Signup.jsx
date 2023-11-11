'use client'

import useSignup from "@/hooks/useSignup"
import { useState } from "react"

export default function Signup({ setIsLogin }) {

  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [confirmPassword, setConfirmPassword] = useState("")

  const { signup } = useSignup()

  async function handleSubmit(e) {
    e.preventDefault()
    await signup(username, password)
    return false
  }

  return (
    <div className="flex items-center h-screen">
      <div className="bg-green-200 rounded-lg max-w-prose m-auto p-8 space-y-4">
        <div className="flex justify-center text-lg font-semibold">
          Sign up
        </div>
        <form onSubmit={handleSubmit}>
          <label>
            <div className="my-2">
              Username
            </div>
            <input
              className="py-1 rounded-lg px-2"
              required
              type="text"
              onChange={(e) => { setUsername(e.target.value) }} value={username}
            />
          </label>
          <label>
            <div className="my-2">
              Password
            </div>
            <input
              className="py-1 rounded-lg px-2"
              required
              type="password"
              onChange={(e) => { setPassword(e.target.value) }} value={password}
            />
          </label>
          <label>
            <div className="my-2">
              Confirm your password
            </div>
            <input
              className="py-1 rounded-lg px-2"
              required
              type="password"
              onChange={(e) => { setConfirmPassword(e.target.value) }} value={confirmPassword} />
          </label>
          <button
          className="block bg-green-900 rounded-lg w-full text-white p-2 mt-4" >
            Create account
          </button>
        </form>
        <div>
          Already have an account?
          <button
          onClick={() => setIsLogin(true)}
          className="block text-blue-700">
            Sign in.
          </button>
        </div>
      </div>
    </div>
  )
}