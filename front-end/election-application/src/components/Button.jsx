import React from 'react'

export default function Button(props) {
  var bgColor = " bg-"+props.bgColor+"-600 "
  var bgColorHover = " hover:bg-"+props.bgColor+"-800 "
  var bgColorActive = " active:bg-"+props.bgColor+"-700 "
  var textColor = " text-"+props.textColor+" "
  var width = " w-full "

  if (props.bgColor == "white") {
    bgColor = " bg-white border-2 "
    bgColorHover = " hover:bg-slate-200 "
    bgColorActive = " active:bg-slate-400 "
    textColor = " text-slate-500"
  }

  const classString = "py-2 rounded-md font-semibold " 
  + bgColor + bgColorHover +  bgColorActive + textColor + width

  return (
    <button className={classString} onClick={props.method}>
      {props.children}
    </button>
  )
}
