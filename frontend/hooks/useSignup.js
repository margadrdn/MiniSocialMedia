import { useRouter } from "next/navigation"
import useCookie from "./useCookie"

export default function useSignup() {

  const router = useRouter()

  const { setCookie } = useCookie()

  const signup = async (username, password) => {
    const response = await fetch('http://127.0.0.1:8080/auth/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password})
    })

    const json = await response.json()

    if (!response.ok) {
      console.log(json.error)
    }

    if (response.ok) {
      setCookie('usertoken', json.token, { expires: new Date(Date.now() + 86400000 )})
      setCookie('username', json.username, { expires: new Date(Date.now() + 86400000 )})
      console.log(document.cookie)
      router.push('/')
    }
  }
  return { signup }
}