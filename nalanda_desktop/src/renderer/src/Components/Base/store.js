import { create } from 'zustand'

export const useStore = create((set) => ({
  token: "",
  user: {}
}))
