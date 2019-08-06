import React  from 'react';
import '../css/Logging.css';
import { useEffect, useState } from 'react';
import {getJwt} from "../jwt";
import {Button, ButtonGroup} from "react-bootstrap";


export const UserList = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        getUsers();
    }, []);

    async function getUsers() {
        let token = getJwt();
        console.log(token);
        const response = await fetch('http://localhost:8080/v1/users',{
            method: 'GET',
            headers:{
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`
            }
        }).catch(error => console.error('Error:', error));
        const users = await response.json();
        console.log(users)
        setUsers(users.users)
    }

    return (
        <div className="Registration">
            {/* main container */}
            <div className="reg">

                {/* Rectangle */}
                <div className="rectangle1">
                    <div className="emptyRectangle1">
                        <h2 className="rightTitle1">User List</h2>
                    </div>
                </div>

                {/*Input form */}
                <div className="reg-form list-users" >
                    <ul>
                        {users.map(user => (
                            <li>ID: {user.id}, First Name: {user.first_name}, Last Name: {user.last_name},
                                UserName: {user.username}</li>
                        ))}
                    </ul>
                </div>

            </div>
            <ButtonGroup className="mr-2 buttns" aria-label="First group">
                <Button href="./adminpage" variant="secondary">Back</Button>
            </ButtonGroup>
        </div>

    );
};

export default UserList;
