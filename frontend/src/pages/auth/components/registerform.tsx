import { TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { Form, Formik } from 'formik';
import * as React from 'react';

interface Values {
    email:      string;
    username:   string;
    password:   string;
    repeatpw:   string;
}

interface Props {
    onSubmit: (values: Values) => void;
}

export const RegisterForm: React.FC<Props> = ({ onSubmit }) => {
    return (
        <Formik 
        initialValues={{email: '', username: '', password: '', repeatpw: '',}} 
        onSubmit={values => {
            onSubmit(values);
        }}
        >
        {({values, handleChange, handleBlur}) => (
            <Form>
                <div>
                <TextField 
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    placeholder="Email"
                    name="email" 
                    value={values.email} 
                    onChange={handleChange}
                    onBlur={handleBlur}
                />
                </div>
                <div>
                <TextField 
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    placeholder="Username"
                    name="username" 
                    value={values.username} 
                    onChange={handleChange}
                    onBlur={handleBlur}
                />
                </div>
                <div>
                <TextField 
                    variant="outlined"
                    margin="normal"
                    required
                    type="password"
                    fullWidth
                    placeholder="Password"
                    name="password" 
                    value={values.password} 
                    onChange={handleChange}
                    onBlur={handleBlur}
                />
                </div>
                <div>
                <TextField 
                    variant="outlined"
                    margin="normal"
                    required
                    type="password"
                    fullWidth
                    placeholder="Repeat Password"
                    name="repeatpw" 
                    value={values.repeatpw} 
                    onChange={handleChange}
                    onBlur={handleBlur}
                />
                </div>
                <Button color="primary" variant="contained" type="submit">Submit</Button>
            </Form>
        )}
        </Formik>
    )
}