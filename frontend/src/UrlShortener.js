import React, { useState } from 'react';
import axios from 'axios';
import { TextField, Button, Typography, Box, Alert, AlertTitle} from '@mui/material';
import FileCopyIcon from '@mui/icons-material/FileCopy';

function UrlShortener() {
  const [longUrl, setLongUrl] = useState('');
  const [shortenedUrl, setShortenedUrl] = useState('');
  const [error, setError] = useState('');
  const [copySuccess, setCopySuccess] = useState('');

  const shortenUrl = async () => {
    try {
      const response = await axios.post('http://localhost:8080/v1/shorten', {
        longurl: longUrl,
      });

      if (response.data.shorturl) {
        setShortenedUrl(response.data.shorturl);
        setError('');
      } else {
        setShortenedUrl('');
        setError('Error shortening URL.');
      }
    } catch (error) {
      console.error('Error:', error);
      const errorMessage = error.response?.data['error'];
      setShortenedUrl('');
      setError(errorMessage || 'An error occurred while connecting to the server.');
    }
  };

  const copyToClipboard = () => {
    const textToCopy = `localhost:8080/v1/${shortenedUrl}`;
    navigator.clipboard.writeText(textToCopy)
      .then(() => setCopySuccess('Copied to clipboard!'))
      .catch(() => setCopySuccess('Copy to clipboard failed.'));
  };

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        height: '100vh',
      }}
    >
      <Typography variant="h4" gutterBottom>
        URL Shortener
      </Typography>
      <TextField
        id="longUrl"
        label="Enter URL"
        variant="outlined"
        value={longUrl}
        onChange={(e) => setLongUrl(e.target.value)}
        sx={{ marginBottom: 2 }}
      />
      <Button variant="contained" onClick={shortenUrl}>
        Shorten!
      </Button>

      {error && (
        <Typography variant="body2" color="error" sx={{ marginTop: 2 }}>
          Error: {error}
        </Typography>
      )}

      {shortenedUrl && (
        <Box sx={{ marginTop: 2 }}>
          <Alert severity="success" sx={{ width: '100%' }}>
            <AlertTitle>Success</AlertTitle>
            Shortened URL created successfully! <strong>localhost:8080/v1/{shortenedUrl}</strong>
          </Alert>
          <Button
            variant="contained"
            onClick={copyToClipboard}
            sx={{ marginTop: 1, textTransform: 'none' }}
            startIcon={<FileCopyIcon />}
          >
            Copy to Clipboard
          </Button>
          <Typography variant="body2" sx={{ marginTop: 1, color: 'green' }}>
            {copySuccess}
          </Typography>
          <Typography variant="body2" sx={{ marginTop: 1 }}>
            Copy and paste the shortened URL in your browser to be redirected.
          </Typography>
        </Box>
      )}
    </Box>
  );
}

export default UrlShortener;
