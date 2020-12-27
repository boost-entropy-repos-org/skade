import * as React from 'react';
import { RegisterForm } from './components/registerform'
import { Grid, Paper, Typography, Avatar, makeStyles } from '@material-ui/core'
import LockIcon from '@material-ui/icons/Lock'

const styles = makeStyles((theme) => ({
    paperStyles: {
        marginTop: theme.spacing(8),
        padding: 50,
        alignItems: "center",
        display: "flex",
        flexDirection: "column"
    },
    gridStyles: {
        minHeight: "70vh",
        alignItems: "center"
    },
}));


const Register = () => {
    const classes = styles()
    return (
        <Grid container justify="center" className={classes.gridStyles}>
            <Paper variant="elevation" elevation={3} className={classes.paperStyles}>
                <Avatar>
                    <LockIcon />
                </Avatar>
                <Typography component="h1" variant="h5">
                    Skade Login
                </Typography>
                <RegisterForm
                    onSubmit={ async ({ email, username, password, repeatpw }) => {
                        console.log(email, username, password, repeatpw)
                        const loginRequest = await fetch("http://localhost:8080/api/v1/createUser/", {
                            method: "POST", 
                        })
                        const loginJson = await loginRequest.json();
                        console.log(loginJson)
                    }}
                />
            </Paper>
        </Grid>
    )
}

export default Register