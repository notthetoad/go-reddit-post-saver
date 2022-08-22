# Ideas for this project

## Make getting posts and comments and inserting them into db concurrent, using channels, sync.Mutex.Lock(), downloading in batches of 100, when done downloading batch, start inserting, start downloading another batch, rinse and repeat
