package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func main() {
	if *Get().File == "" {
		fmt.Printf("Must supply filename\n")
		return
	}
	f, err := os.Open(*Get().File)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	image, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	svc := rekognition.New(session.New())

	/*
		input := &rekognition.DetectLabelsInput{
			Image: &rekognition.Image{
				Bytes: image,
			},
			MaxLabels:     aws.Int64(123),
			MinConfidence: aws.Float64(70.000000),
		}

		result, err := svc.DetectLabels(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case rekognition.ErrCodeInvalidS3ObjectException:
					fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
				case rekognition.ErrCodeInvalidParameterException:
					fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
				case rekognition.ErrCodeImageTooLargeException:
					fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
				case rekognition.ErrCodeAccessDeniedException:
					fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
				case rekognition.ErrCodeInternalServerError:
					fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
				case rekognition.ErrCodeThrottlingException:
					fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
				case rekognition.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case rekognition.ErrCodeInvalidImageFormatException:
					fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
		}
		fmt.Println(result)
	*/

	/*
		input2 := &rekognition.DetectFacesInput{
			Image: &rekognition.Image{
				Bytes: image,
			},
			Attributes: aws.StringSlice([]string{"ALL"}),
		}

		result2, err := svc.DetectFaces(input2)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case rekognition.ErrCodeInvalidS3ObjectException:
					fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
				case rekognition.ErrCodeInvalidParameterException:
					fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
				case rekognition.ErrCodeImageTooLargeException:
					fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
				case rekognition.ErrCodeAccessDeniedException:
					fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
				case rekognition.ErrCodeInternalServerError:
					fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
				case rekognition.ErrCodeThrottlingException:
					fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
				case rekognition.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case rekognition.ErrCodeInvalidImageFormatException:
					fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}
		fmt.Println(result2)
	*/

	/*
		input4 := &rekognition.CreateCollectionInput{
			CollectionId: aws.String("myphotos"),
		}

		result4, err := svc.CreateCollection(input4)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case rekognition.ErrCodeInvalidParameterException:
					fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
				case rekognition.ErrCodeAccessDeniedException:
					fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
				case rekognition.ErrCodeInternalServerError:
					fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
				case rekognition.ErrCodeThrottlingException:
					fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
				case rekognition.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case rekognition.ErrCodeResourceAlreadyExistsException:
					fmt.Println(rekognition.ErrCodeResourceAlreadyExistsException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result4)
	*/

	/*
		input3 := &rekognition.ListFacesInput{
			CollectionId: aws.String("myphotos"),
			MaxResults:   aws.Int64(20),
		}

		result3, err := svc.ListFaces(input3)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case rekognition.ErrCodeInvalidParameterException:
					fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
				case rekognition.ErrCodeAccessDeniedException:
					fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
				case rekognition.ErrCodeInternalServerError:
					fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
				case rekognition.ErrCodeThrottlingException:
					fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
				case rekognition.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case rekognition.ErrCodeInvalidPaginationTokenException:
					fmt.Println(rekognition.ErrCodeInvalidPaginationTokenException, aerr.Error())
				case rekognition.ErrCodeResourceNotFoundException:
					fmt.Println(rekognition.ErrCodeResourceNotFoundException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Println(result3)
	*/

	/*
		input5 := &rekognition.IndexFacesInput{
			CollectionId:    aws.String("myphotos"),
			ExternalImageId: aws.String("eveline"),
			Image: &rekognition.Image{
				Bytes: image,
			},
		}

		result5, err := svc.IndexFaces(input5)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case rekognition.ErrCodeInvalidS3ObjectException:
					fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
				case rekognition.ErrCodeInvalidParameterException:
					fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
				case rekognition.ErrCodeImageTooLargeException:
					fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
				case rekognition.ErrCodeAccessDeniedException:
					fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
				case rekognition.ErrCodeInternalServerError:
					fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
				case rekognition.ErrCodeThrottlingException:
					fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
				case rekognition.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case rekognition.ErrCodeResourceNotFoundException:
					fmt.Println(rekognition.ErrCodeResourceNotFoundException, aerr.Error())
				case rekognition.ErrCodeInvalidImageFormatException:
					fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}
		fmt.Println(result5)
	*/
	/*
		input6 := &rekognition.CompareFacesInput{
			SimilarityThreshold: aws.Float64(90),
			SourceImage: &rekognition.Image{
				Bytes: image,
			},
			TargetImage: &rekognition.Image{
				Bytes: image,
			},
		}

		result6, err := svc.CompareFaces(input6)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case rekognition.ErrCodeInvalidS3ObjectException:
					fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
				case rekognition.ErrCodeInvalidParameterException:
					fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
				case rekognition.ErrCodeImageTooLargeException:
					fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
				case rekognition.ErrCodeAccessDeniedException:
					fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
				case rekognition.ErrCodeInternalServerError:
					fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
				case rekognition.ErrCodeThrottlingException:
					fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
				case rekognition.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case rekognition.ErrCodeResourceNotFoundException:
					fmt.Println(rekognition.ErrCodeResourceNotFoundException, aerr.Error())
				case rekognition.ErrCodeInvalidImageFormatException:
					fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}
		fmt.Println(result6)
	*/

	input7 := &rekognition.SearchFacesByImageInput{
		FaceMatchThreshold: aws.Float64(10),
		CollectionId:       aws.String("myphotos"),
		//MaxFaces:           aws.Int64(10),

		Image: &rekognition.Image{
			Bytes: image,
		},
	}

	result7, err := svc.SearchFacesByImage(input7)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case rekognition.ErrCodeInvalidS3ObjectException:
				fmt.Println(rekognition.ErrCodeInvalidS3ObjectException, aerr.Error())
			case rekognition.ErrCodeInvalidParameterException:
				fmt.Println(rekognition.ErrCodeInvalidParameterException, aerr.Error())
			case rekognition.ErrCodeImageTooLargeException:
				fmt.Println(rekognition.ErrCodeImageTooLargeException, aerr.Error())
			case rekognition.ErrCodeAccessDeniedException:
				fmt.Println(rekognition.ErrCodeAccessDeniedException, aerr.Error())
			case rekognition.ErrCodeInternalServerError:
				fmt.Println(rekognition.ErrCodeInternalServerError, aerr.Error())
			case rekognition.ErrCodeThrottlingException:
				fmt.Println(rekognition.ErrCodeThrottlingException, aerr.Error())
			case rekognition.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(rekognition.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case rekognition.ErrCodeResourceNotFoundException:
				fmt.Println(rekognition.ErrCodeResourceNotFoundException, aerr.Error())
			case rekognition.ErrCodeInvalidImageFormatException:
				fmt.Println(rekognition.ErrCodeInvalidImageFormatException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result7)
}
