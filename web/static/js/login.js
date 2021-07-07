function add_number(num){
    document.pin.pin_field.value += num
}

function delete_number(num){
    let val = document.pin.pin_field.value
    if(val.length < 1){ //i don't know if this is necessary, but it feels better this way
        return
    }
    let x = val.length - 1
    console.log(val.slice(0, x))
    document.pin.pin_field.value = val.slice(0, x)
}