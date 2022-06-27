import React from 'react'
import { Link } from 'react-router-dom'

export default function Navbar(props) {
  return (
    <div className='w-full bg-blue-700 h-12 text-white px-4 uppercase font-semibold flex justify-between'>
      <span className='text-2xl my-auto hover:cursor-pointer hover:text-blue-200'>ELECTION</span>
      <div className='flex my-auto'>
        {props.menu.map((menu) => (
          <Link key={menu.name} to={menu.link}
            className="ml-1 uppercase px-3 py-1 rounded-md
            hover:text-blue-200 active: active:text-blue-100">
            {menu.name}
          </Link>
        ))}

        { props.isLogin ? (
          <div className='my-auto'>
            <Link to="/profile"
              className="ml-1 uppercase px-3 py-1 rounded-md
              hover:text-blue-200 active: active:text-blue-100">
              Profile
            </Link>
            <button onClick={() => console.log("logout")}
              className="ml-1 uppercase px-3 py-1 rounded-md
              hover:text-blue-200 active: active:text-blue-100">
              Logout
            </button>
          </div>
        ) : (
          <div className='my-auto'>
            <Link to="/login"
              className="ml-1 uppercase px-3 py-1 rounded-md
              hover:text-blue-200 active: active:text-blue-100">
              Login
            </Link>
            <Link to="/register"
              className="ml-1 uppercase px-3 py-1 rounded-md
              hover:text-blue-200 active: active:text-blue-100">
              Register
            </Link>
          </div>
        )}
      </div>
    </div>
  )
}
