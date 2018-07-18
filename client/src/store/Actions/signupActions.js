import {
	ROOT_API
} from "./types.js"

export function formOnchange(payload) {
	return (dispatch) => {
		dispatch({
			type: 'SIGNUP_USER',
			payload: payload
		})
	}
}

function clearField(msg) {
	return {
		type: 'CLEAR_FIELD'
	}
}

function errorHandle(err) {
	return {
		type: "HANDLE_ERROR",
		message: err
	}
}

export function SumbitSignUp(payload) {
	return (dispatch) => {
		fetch(`${ROOT_API}/api/admin/user/register`, {
				method: 'POST',
				body: JSON.stringify(payload)
			})
			.then((resp) => resp.json())
			.then(({
				body,
				error
			}) => {
				const respErr = error
				let errMsg
				console.log("=================error", respErr)
				if (respErr === "type request malform") {
					errMsg = 'regiter failed, please field out all field'
					dispatch(errorHandle(errMsg))
				} else if (respErr === "Email already register") {
					errMsg = 'regiter failed, email already register'
					dispatch(errorHandle(errMsg))
				} else {
					dispatch(clearField())
					window.location.href = "/admin";
				}

			}).catch(err => {
				let errMsg = 'regiter failed, please field out all field'
				dispatch(errorHandle(errMsg))
			})
	}
}