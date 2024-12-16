'use client'
import { useRouter } from 'next/navigation'
import Sidebar from '../components/Sidebar'

export default function Home() {
  const router = useRouter()

  return (
    <div>
      <Sidebar/>
    </div>

  );
}
