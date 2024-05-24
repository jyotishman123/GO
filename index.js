function main(){
  try{
    const user = getUser();
    console.log(user)
  } catch (err){
     console.log(err)
  }
} 

function getUser(){
     throw new Error("User not found")
}

main()