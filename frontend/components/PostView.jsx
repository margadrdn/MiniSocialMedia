export default function PostView({ post }) {
  return (
    <div className="bg-green-200 p-4 m-4 rounded-lg max-w-prose">
      <div className="flex gap-4">
        <div className="font-semibold">
          @{post.author}
        </div>
      </div>
      <div className="block">
        {post.body}
      </div>
    </div>
  )
}