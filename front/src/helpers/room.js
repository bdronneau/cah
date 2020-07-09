  import { headRoom } from "@/api/main"

export function roomExist(roomID) {
  return headRoom(roomID).then(response => {
    if (response.status === 200) {
      return true
    }

    return false
  })
}
