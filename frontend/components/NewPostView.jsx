'use client'

import { useRouter } from "next/navigation"
import { useState } from "react"

export default function NewPostView() {
  const router = useRouter()
  const [author, setAuthor] = useState("")
  const [body, setBody] = useState("")
  const [isLoading, setIsLoading] = useState(false)

  const handleSubmit = async (e) => {
    e.preventDefault()
    setIsLoading(true)

    const post = {
      id: "14",
      author,
      body,
    }
    const response = await fetch("http://127.0.0.1:8080/posts", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(post),
    })

    if (response.status === 201) {
      setIsLoading(false)
      setAuthor("")
      setBody("")
      router.refresh()
    }
  }

  return (
    <form onSubmit={handleSubmit} className="bg-green-200 p-4 m-4 rounded-lg max-w-prose">
      <label className="flex gap-4">
        Author
        <input required type="text" onChange={(e) => setAuthor(e.target.value)} value={author} />
      </label>
      <label className="flex gap-4">
        Body
        <input required type="text" onChange={(e) => setBody(e.target.value)} value={body} />
      </label>
      <button className="" disabled={isLoading}>
      {isLoading && "Posting..."}
      {!isLoading && "Post"}
      </button>
    </form>
  )
}