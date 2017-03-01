package main

import(
   "fmt"
   "github.com/aws/aws-sdk-go/aws"
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
   sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
   if err != nil {
      fmt.Println("failed to create session,", err)
    return
   }
   svc := ssm.New(sess)

params := &ssm.AddTagsToResourceInput{
    ResourceId:   aws.String("i-0b13394b7ca44e56f"),             // Required
    ResourceType: aws.String("ManagedInstance"), // Required
    Tags: []*ssm.Tag{ // Required
        { // Required
            Key:   aws.String("SSMXKey"),   // Required
            Value: aws.String("SSMXValue"), // Required
        },
        // More values...
    },
}
resp, err := svc.AddTagsToResource(params)

if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
}

  // Pretty-print the response data.
  fmt.Println(resp)

  fmt.Println("Ok from SSMX")

}
