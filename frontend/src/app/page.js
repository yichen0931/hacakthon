'use client'
import { useRouter } from 'next/navigation'
import Sidebar from '../components/Sidebar'
import OperatingTime from '@/components/OperatingTime'

export default function Home() {
  const router = useRouter()

  return (
    <div className="flex">
      <Sidebar current="Dashboard"/>
      <div className="flex-1 lg:ml-[300px] p-10 overflow-y-auto">
        <OperatingTime/>
      </div>
    </div>
  );
}
