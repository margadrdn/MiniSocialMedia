import { useRouter } from "next/navigation";

export default function useLogout() {
  const router = useRouter()

  const logout = () => {
    document.cookie = `username=; expires=Thu, 01 Jan 1970 00:00:00 GMT`;
    document.cookie = `usertoken=; expires=Thu, 01 Jan 1970 00:00:00 GMT`;

    router.push('/auth')
  }

  return { logout }
}