import useLogin from "@/hooks/useLogin"
import { useState } from "react"

export default function Login({ setIsLogin }) {

  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const { login } = useLogin()

  const handleSubmit = async (e) => {
    e.preventDefault()
    await login(username, password)
    return false
  }

  return (
    <>
      <div className="flex items-center h-screen">
        <div className="bg-green-200 rounded-lg max-w-prose m-auto p-8 space-y-4">
          <div className="flex justify-center text-lg font-semibold">
            Log in
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
            <button className="block bg-green-900 rounded-lg w-full text-white p-2 mt-4">
              Sign in
            </button>
          </form>
          <div>
            Don't have an account?
            <button
              className="block text-blue-700"
              onClick={() => setIsLogin(false)}>
              Create an account.
            </button>
          </div>
        </div>
      </div>
    </>
  )
}