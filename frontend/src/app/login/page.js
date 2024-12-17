"use client"
import logo from '../../assets/foodpanda-app-icon-square.png';
import Image from 'next/image';
import Form from './components/Form';

export default function Login(){
    return (
      <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <Image
            alt="foodpanda"
            src={logo}
            className="mx-auto h-10 w-auto"
          />
          <h2 className="mt-5 text-center text-2xl/9 font-bold tracking-tight text-gray-900 dark:text-white">
            Sign in to your account
          </h2>
        </div>

        <Form/>
      </div>
    )
}