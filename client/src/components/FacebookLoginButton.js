import FacebookLoginBtn from 'react-facebook-login';

function FacebookLoginButton(props) {
  function buttonClicked() {
    console.log("Facebook button clicked.", props)
  }

	function login(response) {
		if (response.status != "unknown") {
      props.login({ 
        username: response.name, 
        picture: response.picture.data.url, 
        userId: response.userID,
      });
		}
	}

	return (
		<FacebookLoginBtn appId = "850650752207328" autoLoad = {false} fields = "name,picture" onClick = {buttonClicked} callback = {login} />
  );
}

export default FacebookLoginButton;