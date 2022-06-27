import React from 'react'

export default function Button(props) {
  var bgColor = ""
  var bgColorHover = ""
  var bgColorActive = ""
  var textColor = ""

  if (props.bgColor === "white") {
    bgColor = " bg-white border-2"
    bgColorHover = " hover:bg-slate-200"
    bgColorActive = " active:bg-slate-400"
    textColor = " text-slate-500 "
  }

  if (props.bgColor === "blue") {
    bgColor = " bg-blue-600"
    bgColorHover = " hover:bg-blue-800"
    bgColorActive = " active:bg-blue-700"
    textColor = " text-white"
  }

  var classString = "py-2 rounded-md font-semibold w-full " + bgColor + bgColorHover +  bgColorActive + textColor

  return (
    <button className={classString} onClick={props.method}>
      {props.children}
    </button>
  )
}
