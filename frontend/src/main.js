import React from "react";
import { render } from "react-dom";
import { BrowserRouter,Route,Switch } from 'react-router-dom';
import { Logging } from "./pages/Logging";
import { MainPage } from "./pages/MainPage";
import { AdminPage } from "./pages/AdminPage";
import { UserList } from "./pages/UserList";
import { PostList } from "./pages/PostList";
import { ModifyUser } from "./pages/ModifyUser";
import { ModifyPost } from "./pages/ModifyPost";
import { DeleteUser } from "./pages/DeleteUser";
import { AdminCreateUser } from "./pages/AdminCreateUser";
import {Registration} from "./pages/Registration";
import {Chat} from "./pages/Chat";
import { Redirect } from "./components/Redirect";
import "./css/bootstrap.min.css";

render(<BrowserRouter>
    <Switch>
        <Route exact path="/" component={Logging}/>
        <Route exact path="/registration" component={Registration}/>


        <Redirect>
            <Route exact path="/chat" component={Chat}/>
            <Route exact path="/mainpage" component={MainPage}/>
            <Route exact path="/adminpage" component={AdminPage}/>
            <Route exact path="/userlist" component={UserList}/>
            <Route exact path="/postlist" component={PostList}/>
            <Route exact path="/usermodify" component={ModifyUser}/>
            <Route exact path="/postmodify" component={ModifyPost}/>
            <Route exact path="/userdelete" component={DeleteUser}/>
            <Route exact path="/usercreate" component={AdminCreateUser}/>
        </Redirect>

    </Switch>
</BrowserRouter>, document.getElementById('app')
);
