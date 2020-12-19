import { TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { Form, Formik } from 'formik';
import * as React from 'react';

interface Values {
    username: string;
    password: string;
}

interface Props {
    onSubmit: (values: Values) => void;
}

export const LoginForm: React.FC<Props> = ({ onSubmit }) => {
    return (
        <Formik 
        initialValues={{username: '', password: ''}} 
        onSubmit={values => {
            onSubmit(values);
        }}
        >
        {({values, handleChange, handleBlur}) => (
            <Form>
                <div>
                <TextField 
                    placeholder="Username"
                    name="username" 
                    value={values.username} 
                    onChange={handleChange}
                    onBlur={handleBlur}
                />
                </div>
                <div>
                <TextField 
                    placeholder="Password"
                    name="password" 
                    value={values.password} 
                    onChange={handleChange}
                    onBlur={handleBlur}
                />
                </div>
                <Button variant="contained" type="submit">Submit</Button>
                <pre>
                    {JSON.stringify(values)}
                </pre>

            </Form>
        )}
        </Formik>
    )
}