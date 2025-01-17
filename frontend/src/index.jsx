/* @refresh reload */
import { render } from 'solid-js/web';
import { Route, Router } from '@solidjs/router';

import './index.css';
import App from './App';
import Dashboard from './pages/dashboard/dashboard';
import SpendingForm from './pages/spendingForm/spendingForm';
import FileManagement from './pages/fileManagement/fileManagement';
import Settings from './pages/settings/settings';

render(() => (
    <Router root={App}>
        <Route path="/" component={Dashboard} />
        <Route path="/add-remove" component={SpendingForm} />
        <Route path="/export" component={FileManagement} />
        <Route path="/settings" component={Settings}/>
    </Router>

), document.getElementById('root'));
