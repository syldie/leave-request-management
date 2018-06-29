let supervisorState = {
    loading: true,
    users: []
}

export default function fetchSupervisorReducer(state = supervisorState, action) {
    switch (action.type) {
        case 'FETCH_LEAVE_PENDING':
            return {
                ...action.payload
            }
        case 'FETCH_LEAVE_ACCEPT':
            return {
                ...action.payload
            }
        case 'FETCH_LEAVE_REJECT':
            return {
                ...action.payload
            }
        case 'ACCEPT_LEAVE_PENDING':
            return {
                ...action.payload
            }
        case 'REJECT_LEAVE_PENDING':
            return {
                ...action.payload
            }
        default:
            return state
    }
}