import { useRouter } from 'next/navigation'
import useCookie from './useCookie'

export default function useLogin() {

  const router = useRouter()

  const { setCookie } = useCookie()

  const login = async (username, password) => {
    const response = await fetch('http://127.0.0.1:8080/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password })
    })

    const json = await response.json()

    if (!response.ok) {
      console.log('error fetching data', response)
    }

    if (response.ok) {
      setCookie('usertoken', json.token, { expires: new Date(Date.now() + 86400000)})
      setCookie('username', json.username, { expires: new Date(Date.now() + 86400000)})
      console.log(document.cookie)
      router.push('/')
    }
  }

  return { login }
}
