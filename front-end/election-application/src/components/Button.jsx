import React from 'react'

export default function Button(props) {
  return (
    <button className='bg-blue-600 px-12 py-2 rounded-md text-white
    hover:bg-blue-800 active:bg-blue-700'>
      {props.children}
    </button>
  )
}
