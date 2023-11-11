import Post from "./Post"

export default async function Posts() {

  const posts = await getPosts()

  return (
    <div>
      {posts.map((post) => <Post post={post} />)}
    </div>
  )
}

async function getPosts() {
  const res = await fetch("http://127.0.0.1:8080/posts", { next: {revalidate: 0 }})
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }
  return res.json()
}