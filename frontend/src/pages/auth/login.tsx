import * as React from 'react';
import { LoginForm } from './components/form'

const Login = () => {
    return (
        <div style={({ textAlign: "center" })}>
            <LoginForm
                onSubmit={({ username, password }) => {
                    console.log(username, password)
                }}
            />
        </div>
    )
}

export default Login